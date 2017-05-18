package {{ .Group }}

import spock.lang.Specification

class SomeClassTest extends Specification {

  def "test increment"() {
    given:
    int input = 1
    int expectedOutput = 2

    when:
    int output = SomeClass.increment(input)

    then:
    output == expectedOutput
  }
}
