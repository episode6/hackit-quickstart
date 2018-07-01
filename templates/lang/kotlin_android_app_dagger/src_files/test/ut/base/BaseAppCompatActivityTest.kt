package {{ .Group }}.base

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
 * Tests [BaseAppCompatActivity]
 */
@RunWith(RobolectricTestRunner::class)
@Config(constants = BuildConfig::class, application = MockspressoTestApp::class)
class BaseAppCompatActivityTest {

  @get:Rule
  val mockspresso = BuildMockspresso.forRobolectric()
      .buildRule()

  @Test
  fun sanityCheck() {
    val controller = Robolectric.buildActivity(TestBaseAppCompatActivity::class.java)
        .setup()
    val activity = controller.get()

    assertThat(activity.supportFragmentInjector()).isNotNull

    controller.pause().stop().destroy()
  }

  companion object {
    class TestBaseAppCompatActivity : BaseAppCompatActivity()
  }
}
