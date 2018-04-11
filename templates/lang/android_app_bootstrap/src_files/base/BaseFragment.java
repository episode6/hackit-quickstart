package {{ .Group }}.base;

import android.app.Activity;
import android.app.Fragment;
import android.content.Context;
import android.os.Build;
import android.os.Bundle;
import android.support.annotation.Nullable;
import android.view.View;

import com.episode6.hackit.android.inject.context.qualifier.ForFragment;
import com.episode6.hackit.disposable.Disposable;
import com.episode6.hackit.disposable.DisposableManager;
import com.episode6.hackit.disposable.butterknife.DisposableButterKnife;
import com.episode6.hackit.pausable.Pausable;
import com.episode6.hackit.pausable.PausableDisposableManager;

import javax.inject.Inject;
import javax.inject.Named;
import javax.inject.Provider;

import dagger.android.AndroidInjection;
import dagger.android.AndroidInjector;
import dagger.android.DispatchingAndroidInjector;
import dagger.android.HasFragmentInjector;

public class BaseFragment extends Fragment implements HasFragmentInjector {

  @Inject DispatchingAndroidInjector<Fragment> mChildFragmentInjector;
  @Inject @ForFragment PausableDisposableManager mDisposables;
  @Inject @Named("forFragmentUi") Provider<DisposableManager> mUiDisposableManagerProvider;

  private DisposableManager mUiDisposables;

  @Override
  public void onAttach(Context context) {
    AndroidInjection.inject(this);
    super.onAttach(context);
  }

  @SuppressWarnings("deprecation")
  @Override
  public void onAttach(Activity activity) {
    if (Build.VERSION.SDK_INT < Build.VERSION_CODES.M) {
      AndroidInjection.inject(this);
    }
    super.onAttach(activity);
  }

  @Override
  public void onViewCreated(View view, @Nullable Bundle savedInstanceState) {
    super.onViewCreated(view, savedInstanceState);
    mUiDisposables = mUiDisposableManagerProvider.get();
    mUiDisposables.addDisposable(DisposableButterKnife.bind(this, view));
  }

  @Override
  public void onResume() {
    super.onResume();
    mDisposables.resume();
  }

  @Override
  public void onPause() {
    mDisposables.pause();
    super.onPause();
  }

  @Override
  public void onDestroyView() {
    if (mUiDisposables != null) {
      mUiDisposables.dispose();
      mUiDisposables = null;
    }
    super.onDestroyView();
  }

  @Override
  public void onDestroy() {
    mDisposables.dispose();
    super.onDestroy();
  }

  @Override
  public AndroidInjector<Fragment> fragmentInjector() {
    return mChildFragmentInjector;
  }

  protected void registerDisposable(Disposable disposable) {
    mDisposables.addDisposable(disposable);
  }

  protected void registerPausable(Pausable pausable) {
    mDisposables.addPausable(pausable);
  }

  protected void registerUiDisposable(Disposable disposable) {
    mUiDisposables.addDisposable(disposable);
  }
}
