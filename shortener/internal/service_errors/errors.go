package service_errors

import "errors"

var ErrShortURLNotFound = errors.New("short url not found")
var ErrInvalidURL = errors.New("shortener: invalid url provided")
var ErrInvalidEncoding = errors.New("shortener: invalid encoding provided")
