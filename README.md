
# server api: api-test.ykypcz.vip
- api-test.ykypcz.vip/api  api接口
- api-test.ykypcz.vip/admin 后台管理

# swagger
- api swagger 文档: http://api-test.ykypcz.vip/api/_swagger
- admin api swagger 文档: http://api-test.ykypcz.vip/admin/_swagger




```nodejs
const fs = require('fs')
const https = require('https')
/**
curl --request POST \
  --url https://api-test.ykypcz.vip/admin/unsafe_api_just_usable_testing/build_test_case \
  --header 'Accept-Language: ' \
  --header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjdXMiOnsidXNlcl9pZCI6IjEwMDAxIn0sImV4cCI6MTczNTA1MDg4Nn0.1omsZu2XwTwEMjiKGR0HyLAwhrvEM_91vBBgDse8-4g' \
  --header 'content-type: application/json' \
  --data '{
    "test_case" : {
        "child": [
            {
                "club_rate": 0.5,
                "invest": 5000,
                "name": "x21"
            }
        ],
        "club_rate": 0.5,
        "invest": 10000,
        "name": "x2"
    },
    "root_user_id": 1
}'
*/
const data = JSON.stringify({
    "test_case" : {
        "child": [
            {
                "club_rate": 0.5,
                "invest": 5000,
                "name": "x21"
            }
        ],
        "club_rate": 0.5,
        "invest": 10000,
        "name": "x2"
    },
    "root_user_id": 1
})

const options = {
  hostname: 'api-test.ykypcz.vip',
  port: 443,
  path: '/admin/unsafe_api_just_usable_testing/build_test_case',
  method: 'POST',
  headers: {
    'Accept-Language': '',
    'Authorization': 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjdXMiOnsidXNlcl9pZCI6IjEwMDAxIn0sImV4cCI6MTczNTA1MDg4Nn0.1omsZu2XwTwEMjiKGR0HyLAwhrvEM_91vBBgDse8-4g',
    'content-type': 'application/json',
    'Content-Length': data.length
  }
}

const req = https.request(options, res => {
  console.log(`statusCode: ${res.statusCode}`)

  res.on('data', d => {
    process.stdout.write(d)
  })
})

req.on('error', error => {
  console.error(error)
})

req.write(data)

req.end()

console.log('send request s
```