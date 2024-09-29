<?php

function isPrivateIP($ip) {
    // Validate the IP address and check if it's in a private range
    return !filter_var(
        $ip, 
        FILTER_VALIDATE_IP, 
        FILTER_FLAG_IPV4 | FILTER_FLAG_NO_PRIV_RANGE |  FILTER_FLAG_NO_RES_RANGE
    );
}

$testIPs = [
    '192.168.1.1',
    '8.8.8.8',
    '10.0.0.1',
    '172.16.0.1',
    '169.254.169.254',
    '127.0.0.1',
    '255.255.255.255', // Broadcast address, should be invalid
];

foreach ($testIPs as $user_ip) {
    if (filter_var($user_ip, FILTER_VALIDATE_IP)) {
        $isPrivate = isPrivateIP($user_ip) ? "true" : "false";
        echo "IP Address: $user_ip | Is Private: $isPrivate\n";
    } else {
        echo "IP Address: $user_ip | Error: Invalid IP\n";
    }
}

?>
