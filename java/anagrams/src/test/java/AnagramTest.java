import static java.nio.charset.StandardCharsets.ISO_8859_1;
import static java.util.Arrays.asList;
import static org.junit.Assert.*;

import java.io.IOException;
import java.net.URISyntaxException;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.List;

import org.junit.Test;

public class AnagramTest {
	private Anagram anagrams;

	@Test
	public void emptyInput(){
		anagrams = new Anagram(new ArrayList<>());
		assertTrue(anagrams.of("").isEmpty());
	}

	@Test
	public void nullInput(){
		anagrams = new Anagram(new ArrayList<>());
		assertTrue(anagrams.of(null).isEmpty());
	}

	@Test
	public void twoLetterInput(){
		anagrams = new Anagram(asList("no", "on", "as", "ops"));
		final List<String> result = anagrams.of("no");

		assertEquals("on", result.get(0));
	}

	@Test
	public void threeLetterInput(){
		anagrams = new Anagram(asList("has", "ash", "had"));
		final List<String> result = anagrams.of("has");

		assertEquals("ash", result.get(0));
	}

	@Test
	public void fourLettersInput(){
		anagrams = new Anagram(asList("cool", "loco", "look"));
		final List<String> result = anagrams.of("cool");

		assertEquals("loco", result.get(0));
	}

	@Test
	public void fiveLettersInput(){
		anagrams = new Anagram(asList("drawer", "redraw", "reward", "warder", "warred"));
		final List<String> result = anagrams.of("drawer");

		assertEquals("redraw", result.get(0));
	}

	@Test
	public void fromWordList() throws IOException, URISyntaxException {
		createFromRealInput();
		final List<String> result = anagrams.of("arrest");

		assertEquals(6, result.size());
	}

	private void createFromRealInput() throws URISyntaxException, IOException {
		final Path path = Paths.get(getClass().getResource("wordlist.txt")
				.toURI());
		anagrams = new Anagram(path, ISO_8859_1);
	}
}