<?php

$user_ip = '::1';
$is_private = filter_var($user_ip, FILTER_VALIDATE_IP, FILTER_FLAG_IPV6 | FILTER_FLAG_NO_PRIV_RANGE) === false ? 'false' : 'true';
$is_loopback = $user_ip === '::1' ? 'true' : 'false';

echo "Language: PHP\n";
echo "IP Address: $user_ip | Is Private: $is_private | Is Loopback: $is_loopback\n";
?>
