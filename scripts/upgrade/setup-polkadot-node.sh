export PARA_HOST=127.0.0.1
cd /home/kien6034/notional/composable-ibc-old/scripts/zombienet # TODO: remove hardfix
process-compose up -f process-compose.yml -t=false & sleep 100