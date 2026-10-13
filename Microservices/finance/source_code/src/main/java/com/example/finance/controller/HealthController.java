package com.example.finance.controller;

import java.lang.management.ManagementFactory;
import java.net.InetAddress;
import java.time.Instant;
import java.util.LinkedHashMap;
import java.util.Map;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class HealthController {
    @GetMapping("/health")
    public Map<String, Object> health() {
        Map<String, Object> body = new LinkedHashMap<>();
        body.put("service", "finance");
        body.put("status", "UP");
        body.put("runtime", "Java " + System.getProperty("java.version"));
        body.put("hostname", hostname());
        body.put("uptimeSeconds", ManagementFactory.getRuntimeMXBean().getUptime() / 1000);
        body.put("timestamp", Instant.now().toString());
        body.put("links", links());
        return body;
    }

    // Optional public URLs of the other services for the home page menu, e.g. AUTH_URL=https://auth.example.com
    private static Map<String, String> links() {
        Map<String, String> links = new LinkedHashMap<>();
        for (String s : new String[] {"auth", "profile", "messenger", "finance", "share"}) {
            String url = System.getenv(s.toUpperCase() + "_URL");
            if (url != null && !url.isEmpty()) {
                links.put(s, url);
            }
        }
        return links;
    }

    private static String hostname() {
        try {
            return InetAddress.getLocalHost().getHostName();
        } catch (Exception e) {
            return "unknown";
        }
    }
}
