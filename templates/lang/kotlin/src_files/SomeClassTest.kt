package {{ .Group }}

import org.junit.Test

import static org.assertj.core.api.Assertions.assertThat

class SomeClassTest {

  @Test
  fun placeholderTest() {
    val input = 1
    val expectedOutput = 2

    val output = SomeClass.increment(input)

    assertThat(output).isEqualTo(expectedOutput)
  }
}
