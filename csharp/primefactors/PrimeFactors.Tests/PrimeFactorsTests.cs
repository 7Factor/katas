using Microsoft.VisualStudio.TestTools.UnitTesting;
using PrimeFactors.Common;

namespace PrimeFactors.Tests
{
    [TestClass]
    public class PrimeFactorsTests
    {
        [TestMethod]
        public void Test1()
        {
            Assert.AreEqual("", string.Join(",", PrimeFactorGenerator.GenerateFactors(1)));
        }

        [TestMethod]
        public void Test2()
        {
            Assert.AreEqual("2", string.Join(",", PrimeFactorGenerator.GenerateFactors(2)));
        }

        [TestMethod]
        public void Test3()
        {
            Assert.AreEqual("3", string.Join(",", PrimeFactorGenerator.GenerateFactors(3)));
        }

        [TestMethod]
        public void Test4()
        {
            Assert.AreEqual("2,2", string.Join(",", PrimeFactorGenerator.GenerateFactors(4)));
        }

        [TestMethod]
        public void Test30()
        {
            Assert.AreEqual("2,3,5", string.Join(",", PrimeFactorGenerator.GenerateFactors(30)));
        }

        [TestMethod]
        public void Test100()
        {
            Assert.AreEqual("2,2,5,5", string.Join(",", PrimeFactorGenerator.GenerateFactors(100)));
        }
    }
}
