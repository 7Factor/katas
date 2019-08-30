import org.junit.Test;
import static org.junit.Assert.*;

public class WrapperTest {

	@Test
	public void emptyString_ShouldReturnEmpty() throws Exception {
		assertEquals("", Wrapper.wrap("", 1));
	}

	@Test
	public void stringShorterThanRowLength_DoesNotWrap() {
		assertEquals("this", Wrapper.wrap("this", 10));
	}

	@Test
	public void stringEqualToRowLength_DoesNotWrap() {
		assertEquals("this", Wrapper.wrap("this", 4));
	}

	@Test
	public void oneWord_ShouldSplit() {
		assertEquals("word\nword", Wrapper.wrap("wordword", 4));
	}

	@Test
	public void oneWord_ShouldSplitMultipleTimes() {
		assertEquals("word\nword\nword", Wrapper.wrap("wordwordword", 4));
	}

	@Test
	public void oneWord_DoesNotThrowOutOfBoundsException() {
		assertEquals("", Wrapper.wrap("", 0));
	}

	@Test
	public void wrapOnWordBoundary() {
		assertEquals("word\nword", Wrapper.wrap("word word", 5));
	}

	@Test
	public void wrapJustBeforeWordBoundary() {
		assertEquals("word\nword", Wrapper.wrap("word word", 4));
	}
}