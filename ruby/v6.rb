require 'ipaddr'

ip = IPAddr.new("::1")
puts "Language: Ruby"
puts "=== Checking IP: #{ip} ==="
puts "Is Private: #{ip.private?}"
puts "Is Loopback: #{ip.loopback?}"
puts
