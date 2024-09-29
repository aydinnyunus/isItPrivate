import ipaddress

def check_ip(ip_str):
    ip = ipaddress.ip_address(ip_str)
    print(f"=== Checking IP: {ip_str} ===")
    print(f"IsPrivate: {ip.is_private}")
    print(f"IsLinkLocal: {ip.is_link_local}")
    print(f"IsLoopback: {ip.is_loopback}")
    print()

def main():
    # Checking 127.0.0.1
    check_ip("127.0.0.1")

    # Checking 169.254.169.254
    check_ip("169.254.169.254")

    # Checking 8.8.8.8
    check_ip("8.8.8.8")

if __name__ == "__main__":
    main()
