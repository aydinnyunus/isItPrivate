import java.net.InetAddress;
import java.net.UnknownHostException;

public class Main {
    public static void main(String[] args) {
        String[] testIPs = {
            "192.168.1.1",
            "8.8.8.8",
            "10.0.0.1",
            "172.16.0.1",
            "169.254.169.254",
            "[::1]",
            "127.0.0.1"
        };

        for (String ip : testIPs) {
            try {
                System.out.printf("IP Address: %-15s | Is Private: %s%n", ip, isPrivate(ip));
            } catch (UnknownHostException e) {
                System.err.printf("IP Address: %-15s | Error: %s%n", ip, e.getMessage());
            }
        }
    }

    public static boolean isPrivate(String ip) throws UnknownHostException {
        InetAddress address = InetAddress.getByName(ip);
        return address.isSiteLocalAddress(); // This will return true for private IPs
    }
}
