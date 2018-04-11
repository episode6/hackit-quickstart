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
import org.robolectric.android.controller.ActivityController;
import org.robolectric.annotation.Config;

import static org.assertj.core.api.Assertions.assertThat;

/**
 * Tests {@link MainActivity}
 */
@RunWith(RobolectricTestRunner.class)
@Config(constants = BuildConfig.class, application = MockspressoTestApp.class)
public class MainActivityTest {

  @Rule public final Mockspresso.Rule mMockspresso = BuildMockspresso.forRobolectric()
      .buildRule();

  @Test
  public void sanityCheck() {
    ActivityController<MainActivity> controller = Robolectric.buildActivity(MainActivity.class).setup();
    MainActivity activity = controller.get();

    assertThat(activity.getTitle())
        .isEqualTo("{{ .CamelName }}");

    controller.pause().stop().destroy();
  }
}
