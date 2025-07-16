#!/bin/bash

echo "📥 Instalando dependencias..."
go mod download
go get github.com/go-gl/cl/v3.0/cl