import json
import requests
import os
from dotenv import load_dotenv


class SlackDriver:

    def __init__(self, webhook_url, channel, username, icon_url):
        self.webhook_url = webhook_url
        self.channel = channel
        self.username = username
        self.icon_url = icon_url

    def send_message(self, message):
        params =  {"text": message, "username": self.username, "channel": self.channel, "icon_url": self.icon_url}
        r = requests.post(self.webhook_url, json.dumps(params))
        print("return ", r.text)


if __name__ == '__main__':
    load_dotenv()
    webhook_url = os.getenv('WEBHOOK_URL')
    channel = os.getenv('CHANNEL')
    icon_url = os.getenv('ICON_URL')
    username = 'Github Actions Bot'
    slack = SlackDriver(webhook_url, channel, username, icon_url)

    # send message
    message = "Hello World! from python"
    slack.send_message(message)
