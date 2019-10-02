import static java.nio.file.Files.lines;
import static java.util.Arrays.asList;
import static java.util.stream.Collectors.toList;

import java.io.IOException;
import java.nio.charset.Charset;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

public class Anagram {
    private final Map<Object, List<String>> dict;

    public Anagram(List<String> words) {
        this.dict = words
                .stream()
                .collect(Collectors.groupingBy(word -> word.length()));
    }

    public Anagram(Path path, Charset charset) throws IOException {
        this(lines(path, charset)
                .map(line -> line.trim())
                .filter(line -> !line.isEmpty())
                .collect(toList()));
    }

    public List<String> of(String input) {
        if (input == null || input.trim().isEmpty())
            return new ArrayList<>();
        return proceed(input);
    }

    private List<String> proceed(String input) {
        return dict.getOrDefault(input.length(), new ArrayList<>())
                .stream()
                .filter(word -> isAnagram(word, input))
                .sorted()
                .collect(toList());
    }

    private boolean isAnagram(String input, String word) {
        return !input.equals(word) && containsSameChars(input, word);
    }

    private boolean containsSameChars(String word1, String word2) {
        return sortLetters(word1).equals(sortLetters(word2));
    }

    private List<String> sortLetters(String input) {
        return asList(input.split(""))
                .stream()
                .sorted()
                .collect(toList());
    }
}