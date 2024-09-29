import ipaddress

ip = ipaddress.ip_address('::1')
print("Language: Python")
print(f"=== Checking IP: {ip} ===")
print(f"Is Private: {ip.is_private}")
print(f"Is Loopback: {ip.is_loopback}")
print()
