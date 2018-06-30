package {{ .Group }}.main

import {{ .Group }}.BuildConfig
import {{ .Group }}.BuildMockspresso
import {{ .Group }}.app.MockspressoTestApp
import org.assertj.core.api.Assertions.assertThat
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.Robolectric
import org.robolectric.RobolectricTestRunner
import org.robolectric.annotation.Config

/**
 * Tests [MainActivity]
 */
@RunWith(RobolectricTestRunner::class)
@Config(constants = BuildConfig::class, application = MockspressoTestApp::class)
class MainActivityTest {

  @get:Rule
  val mockspresso = BuildMockspresso.forRobolectric()
      .buildRule()

  @Test
  fun sanityCheck() {
    val controller = Robolectric.buildActivity(MainActivity::class.java).setup()
    val activity = controller.get()

    assertThat(activity.title)
        .isEqualTo("{{ .CamelName }}")

    controller.pause().stop().destroy()
  }
}
