package {{ .Group }}

import org.assertj.core.api.Assertions.assertThat
import org.junit.Test

class SomeClassTest {

  @Test
  fun placeholderTest() {
    val input = 1
    val expectedOutput = 2

    val output = SomeClass.increment(input)

    assertThat(output).isEqualTo(expectedOutput)
  }
}
