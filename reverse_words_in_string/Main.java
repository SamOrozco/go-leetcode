public class Solution {

    public static void main(String[] args) {
        if (!"blue is sky the".equals(reverseWords("the sky is blue"))) {
            System.out.println("Fail");
            System.exit(1);
        }
    }

    public static String reverseWords(String s) {
        // trim string and split on all white space
        String[] tokens = s.trim().split("\\s+");
        int length = tokens.length;
        StringBuilder bldr = new StringBuilder();
        for (int i = length - 1; i >= 0; i--) {
            if (i != 0) {
                bldr.append(tokens[i]).append(" ");
            } else {
                bldr.append(tokens[i]);
            }
        }

        return bldr.toString();
    }
}