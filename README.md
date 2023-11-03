# VOAH Core Module

![voah-logo-with-text.png](docs/voah-logo-with-text.png)

---

# What is VOAH? - 보아는 어떤 프로젝트인가요?

We were using assistant programs like “Mattermost”, “Notion” and “Github Milestone” to manage on-going projects. However, using various programs resulted in less project continuity and lower workflow efficiency. So we decided to develop an open source project that integrates all the assistant programs. "VOAH"; a messenger, a document manager and a milestone planner in one place.

저희는 기존에 오픈소스 메신저 Mattermost와, 문서 작성을 위한 Notion, Github 마일스톤을 사용하고 있습니다. 그러나 이렇게 분산된 서비스에서 프로젝트를 진행하는 것은 프로젝트의 연속성과 효율성을 저하시키는 원인이 되었습니다. 따라서 메신저, 문서 작성, 마일스톤을 한 곳에서 서로 연계되며 유기적으로 작동할 수 있는 오픈소스 프로젝트 VOAH를 기획하게 되었습니다.

# Quick Start - 빠른 시작

1. Prepare PostgreSQL, Redis and SMTP Server(if you don’t have it, you can use gmail smtp) - PostgreSQL과 Redis, SMTP Server(없다면 다면 gmail smtp를 사용할 수 있습니다)를 준비합니다
   ```bash
   docker network create voah
   ```
   ```yaml
   version: '3'

   services:
     voah-core:
       container_name: voah-core-postgres-dev
       image: postgres:alpine
       restart: always
       hostname: voah-core-postgres
       environment:
         POSTGRES_PASSWORD: password
         POSTGRES_USER: postgres
         POSTGRES_DB: voah-core
       ports:
         - 5432:5432
       expose:
         - 5432
       volumes:
         - ./postgres-data:/var/lib/postgresql/data
     redis:
       container_name: voah-core-redis-dev
       image: redis:alpine
       restart: always
       hostname: voah-core-redis
       ports:
         - 6379:6379
       expose:
         - 6379
       volumes:
         - ./redis-data:/data
       command: redis-server --requirepass password --port 6379
   networks:
     default:
       name: voah
       external: true
   ```

2. Clone the Repository - 레포지토리를 복제합니다

```bash
git clone https://github.com/VOAH-Platform/VOAH-Core.git
```

3. Enter the Directory - 디렉토리로 이동합니다

```bash
cd ./VOAH-Core
```

4. Rename the “setting.example.json” - “setting.example.json” 파일의 이름을 바꿉니다

```bash
mkdir backend
mkdir backend/data
mv ./backend/data/setting.example.json ./backend/data/setting.json
```

5. create “.env” and add environment - “.env” 파일을 만들고 환경변수를 추가합니다

```bash
touch .env

cat << EOF > .env
AUTH_JWT_EXPIRE=3600
AUTH_JWT_SECRET=JWTSECRETasdfiniasejasdf01238
DB_HOST=<<PostgreSQL Host>>
DB_NAME=<<PostgreSQL DB Name>>
DB_PASSWORD=<<PostgreSQL Password>>
DB_PORT=<<PostgreSQL Port>>
DB_USERNAME=<<PostgreSQL Username>>
INSECURESKIPVERIFY=true
REDIS_HOST=<<Redis Host>>
REDIS_LAST_ACTIVITY_DB=1
REDIS_PASSWORD=<<Redis Password>>
REDIS_PORT=<<Redis Port>>
REDIS_SESSION_DB=0
SERVER_CSRF_ORIGIN=*
SERVER_HOST=0.0.0.0
SERVER_HOST_URL=https://voah.company.com
SERVER_PORT=8080
SMTP_HOST=<<SMTP Host>>
SMTP_PASSWORD=<<Mail SMTP Password>>
SMTP_PORT=<<SMTP Port>>
SMTP_STARTTLS=true
SMTP_TLS=false
SMTP_USERNAME=voah@mail.com
SYSTEMADDRESS=voah@mail.com
TZ=Asia/Seoul
VOAH_ROOT_EMAIL=root@company.com
VOAH_ROOT_PW_HASH=<<BCRYPT Hashed Root Password>>
EOF
```

6. Build and Deploy It - 빌드하고 실행하세요

```bash
docker build -t voah-core-back:latest .
docker run --env-file=.env -p 8080:8080 -v ./backend/data:/data --rm -h voah-core --name voah-core voah-core-back:latest
```
Add. Docker Compose - 빌드 없이 실행
   ```yaml
   version: '3'
   
   services:
     voah-core-dev:
       container_name: voah-core-dev
       image: implude/voah-core-dev
       restart: always
       env_file:
         - .env
       expose:
         - 8080
       volumes:
         - ./backend/data:/data
   networks:
     default:
       name: voah
       external: true
   ```
# Customizing - 커스터마이징

"VOAH" support various customization feature. Check below document to see detail.

"보아"는 다양한 사용자 맞춤화 기능을 지원합니다. 아래 문서를 클릭해서 자세히 보실 수 있습니다.
