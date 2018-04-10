package {{ .Group }};

import org.junit.Test;

import static org.assertj.core.api.Assertions.assertThat;

public class SomeClassTest {

  @Test
  public void placeholderTest() {
    int input = 1;
    int expectedOutput = 2;

    int output = SomeClass.increment(input);

    assertThat(output).isEqualTo(expectedOutput);
  }
}
