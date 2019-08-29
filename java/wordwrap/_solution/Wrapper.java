public class Wrapper {

    public static String wrap(String s, int col) {
    if (s.length() <= col)
      return s;
    int space = (s.substring(0, col).lastIndexOf(' '));
    if (space != -1)
      return (s.substring(0, space) + "\n" + wrap(s.substring(space + 1), col));
    else if (s.charAt(col) == ' ')
      return (s.substring(0, col) + "\n" + wrap(s.substring(col + 1), col));
    else
      return (s.substring(0, col) + "\n" + wrap(s.substring(col), col));
  }
}