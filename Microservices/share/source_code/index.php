<?php
$path = parse_url($_SERVER['REQUEST_URI'], PHP_URL_PATH);

// The built-in server keeps no process start time, so read it from /proc on Linux.
function uptime_seconds() {
    $stat = @file_get_contents('/proc/self/stat');
    $sys = @file_get_contents('/proc/uptime');
    if ($stat === false || $sys === false) {
        return null;
    }
    $fields = explode(' ', substr($stat, strrpos($stat, ')') + 2));
    $startTicks = (int) $fields[19];
    return max(0, (int) ((float) $sys - $startTicks / 100));
}

// Optional public URLs of the other services for the home page menu, e.g. AUTH_URL=https://auth.example.com
function links() {
    $links = new stdClass();
    foreach (['auth', 'profile', 'messenger', 'finance', 'share'] as $s) {
        $url = getenv(strtoupper($s) . '_URL');
        if ($url) {
            $links->$s = $url;
        }
    }
    return $links;
}

switch ($path) {
    case '/':
    case '/index.php':
        header('Content-Type: text/html; charset=utf-8');
        readfile(__DIR__ . '/public/index.html');
        break;

    case '/health':
        header('Content-Type: application/json');
        echo json_encode([
            'service' => 'share',
            'status' => 'UP',
            'runtime' => 'PHP ' . PHP_VERSION,
            'hostname' => gethostname(),
            'uptimeSeconds' => uptime_seconds(),
            'timestamp' => gmdate('c'),
            'links' => links(),
        ]);
        break;

    case '/api/hello':
        header('Content-Type: text/plain; charset=utf-8');
        echo "Hello from PHP Share Service";
        break;

    default:
        http_response_code(404);
        header('Content-Type: text/plain; charset=utf-8');
        echo "Not found";
}
