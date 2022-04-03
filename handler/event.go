package handler

import (
	"fmt"
	"crypto/aes"
	"crypto/cipher"
	"io"
	"crypto/rand"
	"encoding/base64"
	"github.com/amirkhaki/cnbzdtr/entity"
	"github.com/amirkhaki/cnbzdtr/config"
	dg "github.com/bwmarrin/discordgo"
)

// score change event handler
type SEH struct {
	s    *dg.Session
	lvls *entity.Levels
	cfg config.Config
}

func (sh *SEH) Handle(u *entity.User) error {
	crrntLevel := sh.lvls.Level(u.MostScore)
	prevLevel := sh.lvls.Level(u.PrevMostScore)
	if crrntLevel == prevLevel {
		return nil
	}
	ch, err := sh.s.UserChannelCreate(u.ID)
	if err != nil {
		return fmt.Errorf("Could not create DM channel with user: %w", err)
	}
	uidEnc, err := encrypt([]byte(sh.cfg.Cipher_key), u.ID)
	if err != nil {
		return fmt.Errorf("Could not encrypt user id: %w", err)
	}
	message := fmt.Sprintf("%s?user=%s", crrntLevel.Url, uidEnc)
	_, err = sh.s.ChannelMessageSend(ch.ID, message)
	if err != nil {
		return fmt.Errorf("Could not send DM to user: %w", err)
	}
	return nil
}


func encrypt(key []byte, message string) (encmess string, err error) {
	plainText := []byte(message)

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	//returns to base64 encoded string
	encmess = base64.URLEncoding.EncodeToString(cipherText)
	return
}
