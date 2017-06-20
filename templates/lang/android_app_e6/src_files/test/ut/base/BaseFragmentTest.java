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
import org.robolectric.android.controller.FragmentController;
import org.robolectric.annotation.Config;

import static org.fest.assertions.api.Assertions.assertThat;

/**
 * Tests {@link BaseFragment}
 */
@RunWith(RobolectricTestRunner.class)
@Config(constants = BuildConfig.class, application = MockspressoTestApp.class)
public class BaseFragmentTest {

  @Rule public final Mockspresso.Rule mMockspresso = BuildMockspresso.forRobolectric()
      .buildRule();

  @Test
  public void sanityCheck() {
    FragmentController<BaseFragment> controller = Robolectric.buildFragment(BaseFragment.class)
        .create()
        .start()
        .resume()
        .visible();
    BaseFragment fragment = controller.get();

    assertThat(fragment.fragmentInjector()).isNotNull();

    controller.pause().stop().destroy();
  }
}
