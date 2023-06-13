# IPSW Go
Automatically download the latest apple firmwares you subscribed.

Developed by 0xJacky

Data source comes from betahub.cn

## Usage

**IMPORTANT**: Replace `/mnt/user/downloads` with your host path.

```
docker run -dit \
  --name=ipsw-go \
  --restart=always \
  -e Workers=4 \
  -e Identifiers=iPhone13,1:iPhone13,2:iPad8,9 \
  -e TZ=Asia/Shanghai \
  -e CheckAt=02:00,04:00,06:00 \
  -e LastTwoVer=true \
  -v /mnt/user/downloads:/downloads \
  uozi/ipsw-go:latest
```

## Environment Variable
- **Workers**: (optional) concurrent download workers num, default: 1.
- **Identifiers**: (required) iDevice identifiers, use `:` to separate.
- **TZ**: (required) timezone string.
- **CheckAt**: (required) set time to run, use `,` to separate.
- **LastTwoVer**: (optional) download last two version firmwares.
