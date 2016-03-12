// javac HelloWorld.java
// java HelloWorld

public class HelloWorld {
    public static void main(String[] args) throws Exception {
        System.out.println("Hello World");
        while (true) {
            System.out.println("STDOUT of a sub-shell");
            Thread.sleep(5000);
        }
    }
}