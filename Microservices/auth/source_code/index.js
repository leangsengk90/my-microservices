const http = require('http');
const fs = require('fs');
const os = require('os');
const path = require('path');

const PORT = 3000;
const page = fs.readFileSync(path.join(__dirname, 'public', 'index.html'));

// Optional public URLs of the other services for the home page menu, e.g. PROFILE_URL=https://profile.example.com
const SERVICES = ['auth', 'profile', 'messenger', 'finance', 'share'];
const links = Object.fromEntries(
  SERVICES.filter(s => process.env[`${s.toUpperCase()}_URL`]).map(s => [s, process.env[`${s.toUpperCase()}_URL`]])
);

http.createServer((req, res) => {
  const url = req.url.split('?')[0];

  if (url === '/' || url === '/index.html') {
    res.writeHead(200, { 'Content-Type': 'text/html; charset=utf-8' });
    res.end(page);
  } else if (url === '/health') {
    res.writeHead(200, { 'Content-Type': 'application/json' });
    res.end(JSON.stringify({
      service: 'auth',
      status: 'UP',
      version: process.env.APP_VERSION || 'dev',
      runtime: `Node.js ${process.version}`,
      hostname: os.hostname(),
      uptimeSeconds: Math.floor(process.uptime()),
      timestamp: new Date().toISOString(),
      links,
    }));
  } else if (url === '/api/hello') {
    res.writeHead(200, { 'Content-Type': 'text/plain; charset=utf-8' });
    res.end('Hello from Node.js Auth Service');
  } else {
    res.writeHead(404, { 'Content-Type': 'text/plain; charset=utf-8' });
    res.end('Not found');
  }
}).listen(PORT, () => {
  console.log(`Auth service running on ${PORT}`);
});
