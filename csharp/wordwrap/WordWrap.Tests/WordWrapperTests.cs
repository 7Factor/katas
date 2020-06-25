using Microsoft.VisualStudio.TestTools.UnitTesting;
using WordWrap.Common;

namespace WordWrap.Tests
{
    [TestClass]
    public class WordWrapperTests
    {
        [TestMethod]
        public void WrapWords_WhenPassedEmptyString_ShouldReturnEmptyString()
        {
            Assert.AreEqual(string.Empty, WordWrapper.WrapWords(string.Empty, 1));
        }

        [TestMethod]
        public void WrapWords_WhenPassedStringShorterThanRowLength_ShouldReturnInputString()
        {
            Assert.AreEqual("Word", WordWrapper.WrapWords("Word", 5));
        }

        [TestMethod]
        public void WrapWords_WhenPassedStringEqualToRowLength_ShouldReturnInputString()
        {
            Assert.AreEqual("Word", WordWrapper.WrapWords("Word", 4));
        }

        [TestMethod]
        public void WrapWords_WhenPassedStringWithNoSpacesGreaterThanRowLength_ShouldReturnStringSplitAtRowLength()
        {
            Assert.AreEqual("WordW\nord", WordWrapper.WrapWords("WordWord", 5));
        }

        [TestMethod]
        public void WrapWords_WhenPassedLongerStringWithNoSpacesGreaterThanRowLength_ShouldReturnStringSplitAtRowLength()
        {
            Assert.AreEqual("WordW\nordWo\nrd", WordWrapper.WrapWords("WordWordWord", 5));
        }

        [TestMethod]
        public void WrapWords_WhenPassedRowLengthLessThanOne_ShouldReturnInputString()
        {
            Assert.AreEqual("Word", WordWrapper.WrapWords("Word", 0));
        }

        [TestMethod]
        public void WrapWords_WhenPassedStringWithSpacesAtRowLength_ShouldReturnStringSplitAtSpaces()
        {
            Assert.AreEqual("Word\nWord", WordWrapper.WrapWords("Word Word", 5));
        }

        [TestMethod]
        public void WrapWords_WhenPassedStringWithSpacesAfterRowLength_ShouldReturnStringSplitAtSpaces()
        {
            Assert.AreEqual("Word\nWord", WordWrapper.WrapWords("Word Word", 4));
        }

        [TestMethod]
        public void WrapWords_WhenPassedComplexString_ShouldReturnStringCorrectly()
        {
            Assert.AreEqual("Word\nWordWord\nWord", WordWrapper.WrapWords("Word WordWord Word", 8));
        }

        [TestMethod]
        public void WrapWords_WhenPassedComplexString_ShouldReturnStringCorrectlyTwo()
        {
            Assert.AreEqual("WordWord\nWord\nWordWord\nWord", WordWrapper.WrapWords("WordWordWord WordWord Word", 8));
        }
    }
}
