# 確認事項
- gorm
```
% curl http://localhost:8080/users/20                    
{"ID":20,"CreatedAt":"2022-09-06T22:04:10.625689+09:00","UpdatedAt":"2022-09-06T22:04:10.625689+09:00","DeletedAt":null,"Name":"Kyra Green","GroupID":3,"Group":{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"Name":""}}
% curl http://localhost:8080/users/21
% curl -X POST http://localhost:8080/users -d 'name=hoge'
{"ID":21,"CreatedAt":"2022-09-06T22:04:52.111813+09:00","UpdatedAt":"2022-09-06T22:04:52.111813+09:00","DeletedAt":null,"Name":"hoge","GroupID":0,"Group":{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"Name":""}}
% curl http://localhost:8080/users/21             
{"ID":21,"CreatedAt":"2022-09-06T22:04:52.111813+09:00","UpdatedAt":"2022-09-06T22:04:52.111813+09:00","DeletedAt":null,"Name":"hoge","GroupID":0,"Group":{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"Name":""}}
% curl -X PUT http://localhost:8080/users/21 -d 'name=hoge'
{"ID":21,"CreatedAt":"2022-09-06T22:04:52.111813+09:00","UpdatedAt":"2022-09-06T22:05:14.689384+09:00","DeletedAt":null,"Name":"hoge","GroupID":0,"Group":{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"Name":""}}
% curl http://localhost:8080/users/21                      
{"ID":21,"CreatedAt":"2022-09-06T22:04:52.111813+09:00","UpdatedAt":"2022-09-06T22:05:14.689384+09:00","DeletedAt":null,"Name":"hoge","GroupID":0,"Group":{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"Name":""}}
% curl -X DELETE http://localhost:8080/users/21   
deleted% 
```
