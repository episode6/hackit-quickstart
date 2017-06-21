package {{ .Group }}.base;

import android.os.Bundle;
import android.support.annotation.Nullable;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;

import com.episode6.hackit.android.inject.context.qualifier.ForFragment;
import com.episode6.hackit.disposable.Disposable;
import com.episode6.hackit.disposable.DisposableManager;
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
import org.robolectric.android.controller.FragmentController;
import org.robolectric.annotation.Config;

import javax.inject.Named;

import static org.fest.assertions.api.Assertions.assertThat;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.inOrder;
import static org.mockito.Mockito.verifyNoMoreInteractions;

/**
 * Tests {@link BaseFragment}
 */
@RunWith(RobolectricTestRunner.class)
@Config(constants = BuildConfig.class, application = MockspressoTestApp.class)
public class BaseFragmentTest {

  @Rule public final Mockspresso.Rule mMockspresso = BuildMockspresso.forRobolectric()
      .buildRule();

  @Mock @ForFragment PausableDisposableManager mDisposables;
  @Mock @Named("forFragmentUi") DisposableManager mUiDisposables;

  @Mock @Unmapped Disposable mDisposable;
  @Mock @Unmapped Pausable mPausable;

  @Test
  public void sanityCheck() {
    InOrder inOrder = inOrder(mDisposables, mUiDisposables);
    FragmentController<BaseFragment> controller = Robolectric.buildFragment(BaseFragment.class)
        .create()
        .start()
        .resume()
        .visible();
    BaseFragment fragment = controller.get();

    inOrder.verify(mDisposables).resume();
    assertThat(fragment.fragmentInjector()).isNotNull();

    fragment.registerDisposable(mDisposable);
    inOrder.verify(mDisposables).addDisposable(mDisposable);

    fragment.registerPausable(mPausable);
    inOrder.verify(mDisposables).addPausable(mPausable);

    controller.pause();
    inOrder.verify(mDisposables).pause();

    controller.stop().destroy();
    inOrder.verify(mDisposables).dispose();
    inOrder.verifyNoMoreInteractions();
    verifyNoMoreInteractions(mUiDisposables);
  }

  @Test
  public void sanityCheckWithView() {
    InOrder inOrder = inOrder(mDisposables, mUiDisposables);
    FragmentController<BaseFragmentWithView> controller = Robolectric.buildFragment(BaseFragmentWithView.class)
        .create()
        .start()
        .resume()
        .visible();
    BaseFragment fragment = controller.get();

    inOrder.verify(mUiDisposables).addDisposable(any(Disposable.class)); // butterknife
    inOrder.verify(mDisposables).resume();
    assertThat(fragment.fragmentInjector()).isNotNull();

    fragment.registerUiDisposable(mDisposable);
    inOrder.verify(mUiDisposables).addDisposable(mDisposable);

    controller.pause();
    inOrder.verify(mDisposables).pause();

    controller.stop().destroy();
    inOrder.verify(mUiDisposables).dispose();
    inOrder.verify(mDisposables).dispose();
    inOrder.verifyNoMoreInteractions();
  }

  public static class BaseFragmentWithView extends BaseFragment {
    @Nullable
    @Override
    public View onCreateView(LayoutInflater inflater, @Nullable ViewGroup container, Bundle savedInstanceState) {
      return new View(container.getContext());
    }
  }
}
