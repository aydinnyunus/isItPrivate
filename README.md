# isItPrivate

This project explores the classification of IP addresses across various programming languages, specifically focusing on private and loopback addresses. The goal is to analyze the inconsistencies in how different languages treat these IPs and highlight potential security implications, especially in cloud environments.

## Overview

The project examines several IP addresses, including:

- **Loopback IP Address**: `127.0.0.1` and `::1`
- **Link-Local Address**: `169.254.169.254`
- **Private IP Addresses**: Various private IPs (e.g., `192.168.1.1`, `10.0.0.1`)

## Languages Covered

The analysis includes implementations in the following programming languages:

- Go
- Java
- Node.js
- PHP
- Python
- Ruby

## Security Implications

Inconsistent classification of IP addresses can lead to vulnerabilities such as Server-Side Request Forgery (SSRF) in cloud environments like AWS and Google Cloud. The project aims to raise awareness of these issues and promote standardization in IP classification.

## Getting Started

To explore the code and findings, clone the repository:

```bash
git clone https://github.com/aydinnyunus/isItPrivate.git
cd isItPrivate
```

Each language's implementation is located in its respective folder. You can run the scripts to see the output for the specified IP addresses.

## Contributing

Contributions are welcome! If you have suggestions or improvements, feel free to open an issue or submit a pull request.

## Contact

[<img target="_blank" src="https://img.icons8.com/bubbles/100/000000/linkedin.png" title="LinkedIn">](https://linkedin.com/in/yunus-ayd%C4%B1n-b9b01a18a/) [<img target="_blank" src="https://img.icons8.com/bubbles/100/000000/github.png" title="Github">](https://github.com/aydinnyunus/WhatsappBOT) [<img target="_blank" src="https://img.icons8.com/bubbles/100/000000/instagram-new.png" title="Instagram">](https://instagram.com/aydinyunus_/) [<img target="_blank" src="https://img.icons8.com/bubbles/100/000000/twitter-squared.png" title="LinkedIn">](https://twitter.com/aydinnyunuss)
