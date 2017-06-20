package {{ .Group }}.base;

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

import static org.fest.assertions.api.Assertions.assertThat;

/**
 * Tests {@link BaseAppCompatActivity}
 */
@RunWith(RobolectricTestRunner.class)
@Config(constants = BuildConfig.class, application = MockspressoTestApp.class)
public class BaseAppCompatActivityTest {

  @Rule public final Mockspresso.Rule mMockspresso = BuildMockspresso.forRobolectric()
      .buildRule();

  @Test
  public void sanityCheck() {
    ActivityController<BaseAppCompatActivity> controller = Robolectric.buildActivity(BaseAppCompatActivity.class)
        .setup();
    BaseAppCompatActivity activity = controller.get();

    assertThat(activity.fragmentInjector()).isNotNull();

    controller.pause().stop().destroy();
  }
}
