ifconfig eth0:1 172.29.1.1/1
ifconfig eth0:2 172.29.1.2/1
ifconfig eth0:3 172.29.1.3/1
ifconfig eth0:4 172.29.1.4/1
ifconfig eth0:5 172.29.1.5/1
ifconfig eth0:6 172.29.1.6/1

ifconfig eth0:7 172.29.2.1/1
ifconfig eth0:8 172.29.2.2/1
echo "Finsh ifconfig eth0:1~8 config"

export RABBITMQ_SERVER=amqp://zhengyu:741213@172.29.158.120:5672
# export RABBITMQ_SERVER=amqp://zhengyu:741213@3r48j22500.yicp.fun:14472

export ES_SERVER=172.29.158.120:9200
# export ES_SERVER=3r48j22500.yicp.fun:33715

echo "Finsh RABBITMQ_SERVER config:" $RABBITMQ_SERVER
echo "Finsh ES_SERVER config:" $ES_SERVER
