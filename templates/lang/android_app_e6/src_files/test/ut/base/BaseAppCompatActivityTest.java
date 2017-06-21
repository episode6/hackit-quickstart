package {{ .Group }}.base;

import com.episode6.hackit.android.inject.context.qualifier.ForContextScope;
import com.episode6.hackit.disposable.Disposable;
import com.episode6.hackit.mockspresso.Mockspresso;
import com.episode6.hackit.mockspresso.annotation.Unmapped;
import com.episode6.hackit.pausable.Pausable;
import com.episode6.hackit.pausable.PausableDisposableManager;

import {{ .Group }}.BuildConfig;
import {{ .Group }}.BuildMockspresso;
import {{ .Group }}.app.MockspressoTestApp;

import org.junit.Rule;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.mockito.InOrder;
import org.mockito.Mock;
import org.robolectric.Robolectric;
import org.robolectric.RobolectricTestRunner;
import org.robolectric.android.controller.ActivityController;
import org.robolectric.annotation.Config;

import static org.fest.assertions.api.Assertions.assertThat;
import static org.mockito.Mockito.inOrder;

/**
 * Tests {@link BaseAppCompatActivity}
 */
@RunWith(RobolectricTestRunner.class)
@Config(constants = BuildConfig.class, application = MockspressoTestApp.class)
public class BaseAppCompatActivityTest {

  @Rule public final Mockspresso.Rule mMockspresso = BuildMockspresso.forRobolectric()
      .buildRule();

  @Mock @ForContextScope PausableDisposableManager mDisposables;

  @Mock @Unmapped Disposable mDisposable;
  @Mock @Unmapped Pausable mPausable;

  @Test
  public void sanityCheck() {
    InOrder inOrder = inOrder(mDisposables);

    ActivityController<BaseAppCompatActivity> controller = Robolectric.buildActivity(BaseAppCompatActivity.class)
        .setup();
    BaseAppCompatActivity activity = controller.get();

    inOrder.verify(mDisposables).resume();
    assertThat(activity.fragmentInjector()).isNotNull();

    activity.registerDisposable(mDisposable);
    inOrder.verify(mDisposables).addDisposable(mDisposable);

    activity.registerPausable(mPausable);
    inOrder.verify(mDisposables).addPausable(mPausable);

    controller.pause();
    inOrder.verify(mDisposables).pause();

    controller.stop().destroy();
    inOrder.verify(mDisposables).dispose();
  }
}
