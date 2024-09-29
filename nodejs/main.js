const ip = require('ip');

function checkIp(ipStr) {
    console.log(`=== Checking IP: ${ipStr} ===`);
    console.log(`IsPrivate: ${ip.isPrivate(ipStr)}`);
    console.log(`IsLoopback: ${ip.isLoopback(ipStr)}`);
    console.log();
}

function main() {
    // Checking 127.0.0.1
    checkIp("127.0.0.1");

    // Checking 169.254.169.254
    checkIp("169.254.169.254");

    // Checking 8.8.8.8
    checkIp("8.8.8.8");
}

main();
