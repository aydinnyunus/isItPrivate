require 'ipaddr'

# Define the IP addresses to check
test_ips = [
  "192.168.1.1",
  "8.8.8.8",
  "169.254.169.254",
  "10.0.0.1",
  "172.16.0.1",
  "127.0.0.1",
  "255.255.255.255"
]


# Check each IP address and print the result
test_ips.each do |ip|
  is_private = IPAddr.new(ip).private?
  puts "IP Address: #{ip.ljust(15)} | Is Private: #{is_private}"
end
