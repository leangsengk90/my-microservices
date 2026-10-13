import os
import platform
import socket
import time
from datetime import datetime, timezone

from flask import Flask, jsonify

app = Flask(__name__)
START_TIME = time.time()

# Optional public URLs of the other services for the home page menu, e.g. AUTH_URL=https://auth.example.com
SERVICES = ['auth', 'profile', 'messenger', 'finance', 'share']
LINKS = {s: os.environ[f'{s.upper()}_URL'] for s in SERVICES if os.environ.get(f'{s.upper()}_URL')}


@app.route('/')
def home():
    return app.send_static_file('index.html')


@app.route('/health')
def health():
    return jsonify(
        service='profile',
        status='UP',
        version=os.environ.get('APP_VERSION') or 'dev',
        runtime=f'Python {platform.python_version()}',
        hostname=socket.gethostname(),
        uptimeSeconds=int(time.time() - START_TIME),
        timestamp=datetime.now(timezone.utc).isoformat(),
        links=LINKS,
    )


@app.route('/api/hello')
def hello():
    return "Hello from Python Flask Profile Service"


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=3001)
