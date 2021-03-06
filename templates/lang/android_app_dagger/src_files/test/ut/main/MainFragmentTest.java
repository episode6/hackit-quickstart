package {{ .Group }}.main;

import com.episode6.hackit.mockspresso.Mockspresso;

import {{ .Group }}.BuildConfig;
import {{ .Group }}.BuildMockspresso;
import {{ .Group }}.app.MockspressoTestApp;

import org.junit.Rule;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.robolectric.Robolectric;
import org.robolectric.RobolectricTestRunner;
import org.robolectric.annotation.Config;
import org.robolectric.shadows.support.v4.SupportFragmentController;

import static org.assertj.core.api.Assertions.assertThat;

/**
 * Test {@link MainFragment}
 */
@RunWith(RobolectricTestRunner.class)
@Config(constants = BuildConfig.class, application = MockspressoTestApp.class)
public class MainFragmentTest {

  @Rule public final Mockspresso.Rule mMockspresso = BuildMockspresso.forRobolectric()
      .buildRule();

  @Test
  public void sanityCheck() {
    SupportFragmentController<MainFragment> controller = SupportFragmentController.of(new MainFragment())
        .create()
        .start()
        .resume()
        .visible();

    MainFragment fragment = controller.get();
    assertThat(fragment.mTextView.getText()).isEqualTo("Hello Earth!");

    controller.pause().stop().destroy();
  }
}
