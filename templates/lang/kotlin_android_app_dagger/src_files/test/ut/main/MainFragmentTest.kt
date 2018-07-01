package {{ .Group }}.main

import com.episode6.hackit.mockspresso.Mockspresso
import {{ .Group }}.BuildConfig
import {{ .Group }}.BuildMockspresso
import {{ .Group }}.app.MockspressoTestApp
import kotlinx.android.synthetic.main.fragment_main.*
import org.assertj.core.api.Assertions.assertThat
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.RobolectricTestRunner
import org.robolectric.annotation.Config
import org.robolectric.shadows.support.v4.SupportFragmentController

/**
 * Test [MainFragment]
 */
@RunWith(RobolectricTestRunner::class)
@Config(constants = BuildConfig::class, application = MockspressoTestApp::class)
class MainFragmentTest {

  @get:Rule
  val mockspresso = BuildMockspresso.forRobolectric()
      .buildRule()

  @Test
  fun sanityCheck() {
    val controller = SupportFragmentController.of(MainFragment())
        .create()
        .start()
        .resume()
        .visible()

    val fragment = controller.get()
    assertThat(fragment.tv_hello_world.text).isEqualTo("Hello Earth!")

    controller.pause().stop().destroy()
  }
}
