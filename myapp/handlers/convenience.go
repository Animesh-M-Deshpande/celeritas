package handlers

import (
	"context"
	"net/http"

	"github.com/Animesh-M-Deshpande/celeritas"
)

func (h *Handlers) render(w http.ResponseWriter, r *http.Request, tmpl string, variables, data interface{}) error {

	return h.App.Render.Page(w, r, tmpl, variables, data)
}

func (h *Handlers) sessionPut(context context.Context, key string, val interface{}) {
	h.App.Session.Put(context, key, val)
}

func (h *Handlers) sessionHas(context context.Context, key string) bool {
	return h.App.Session.Exists(context, key)
}

func (h *Handlers) sessionGet(context context.Context, key string) interface{} {
	return h.App.Session.Get(context, key)
}

func (h *Handlers) sessionRemove(context context.Context, key string) {
	h.App.Session.Remove(context, key)
}

func (h *Handlers) sessionRenew(ctx context.Context) error {
	return h.App.Session.RenewToken(ctx)
}

func (h *Handlers) sessionDestroy(ctx context.Context) error {
	return h.App.Session.Destroy(ctx)
}

func (h *Handlers) randomString(n int) string {
	return h.App.RandomString(n)
}

func (h *Handlers) Encrypt(text string) (string, error) {
	enc := celeritas.Encryption{Key: []byte(h.App.EncryptionKey)}
	encrypted, err := enc.Encrypt(text)

	if err != nil {
		return "", err
	}

	return encrypted, nil
}

func (h *Handlers) Decrypt(crypto string) (string, error) {
	enc := celeritas.Encryption{Key: []byte(h.App.EncryptionKey)}
	text, err := enc.Decrypt(crypto)

	if err != nil {
		return "", err
	}

	return text, nil
}
