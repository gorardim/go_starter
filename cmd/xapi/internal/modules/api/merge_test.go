package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeOpenApiYamlContent(t *testing.T) {
	src := `
openapi: 3.0.0
info:
  title: Append
  version: 1.0.0
paths:
  /api/upload/image:
    post:
      description: "图片上传"
      summary: "图片上传"
      parameters:
        - in: header
          name: Authorization
          description: Authorization token
          required: true
          schema:
            type: string
      requestBody:
        content:
          multipart/form-data:
            schema:
              type: object
              properties:
                file:
                  type: string
                  format: binary
      responses:
        200:
          description: OK
          content:
            application/json:
              schema:
                type: object
                properties:
                  url:
                    type: string
                    description: 图片地址
`
	dst := `paths:
  /api/upload/image2:
    post:
      description: "图片上传"
      summary: "图片上传"
      parameters:
        - in: header
          name: Authorization
          description: Authorization token
          required: true
          schema:
            type: string
      requestBody:
        content:
          multipart/form-data:
            schema:
              type: object
              properties:
                file:
                  type: string
                  format: binary
      responses:
        200:
          description: OK
          content:
            application/json:
              schema:
                type: object
                properties:
                  url:
                    type: string
                    description: 图片地址`
	content, err := MergeOpenApiYamlContent([]byte(src), []byte(dst))
	assert.NoError(t, err)
	t.Log(content)
}
