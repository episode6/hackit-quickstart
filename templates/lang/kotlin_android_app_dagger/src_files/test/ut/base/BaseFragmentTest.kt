package {{ .Group }}.base

import {{ .Group }}.BuildConfig
import {{ .Group }}.BuildMockspresso
import {{ .Group }}.app.MockspressoTestApp
import org.assertj.core.api.Assertions.assertThat
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.RobolectricTestRunner
import org.robolectric.annotation.Config
import org.robolectric.shadows.support.v4.SupportFragmentController

/**
 * Tests [BaseFragment]
 */
@RunWith(RobolectricTestRunner::class)
@Config(constants = BuildConfig::class, application = MockspressoTestApp::class)
class BaseFragmentTest {

  @get:Rule
  val mockspresso = BuildMockspresso.forRobolectric()
      .buildRule()

  @Test
  fun sanityCheck() {
    val controller = SupportFragmentController.of(TestBaseFragment())
        .create()
        .start()
        .resume()
        .visible()
    val fragment = controller.get()

    assertThat(fragment.supportFragmentInjector()).isNotNull()

    controller.pause().stop().destroy()
  }

  companion object {
    class TestBaseFragment : BaseFragment()
  }
}
