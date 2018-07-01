package {{ .Group }};

import org.assertj.core.api.Assertions.assertThat
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.Robolectric
import org.robolectric.RobolectricTestRunner
import org.robolectric.annotation.Config

@RunWith(RobolectricTestRunner::class)
@Config(constants = BuildConfig::class)
class MainActivityTest {

  @Test
  fun placeholderTest() {
    val activity = Robolectric.setupActivity(MainActivity::class.java)

    assertThat(activity.title)
        .isEqualTo("{{ .CamelName }}")
  }
}
