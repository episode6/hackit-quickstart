package {{ .Group }}.base;

import android.app.Fragment;
import android.os.Bundle;
import android.support.annotation.Nullable;
import android.support.v7.app.AppCompatActivity;

import com.episode6.hackit.android.inject.context.qualifier.ForContextScope;
import com.episode6.hackit.disposable.Disposable;
import com.episode6.hackit.pausable.Pausable;
import com.episode6.hackit.pausable.PausableDisposableManager;

import javax.inject.Inject;

import dagger.android.AndroidInjection;
import dagger.android.AndroidInjector;
import dagger.android.DispatchingAndroidInjector;
import dagger.android.HasFragmentInjector;

public class BaseAppCompatActivity extends AppCompatActivity implements HasFragmentInjector {

  @Inject DispatchingAndroidInjector<Fragment> mFragmentInjector;
  @Inject @ForContextScope PausableDisposableManager mDisposables;

  @Override
  protected void onCreate(@Nullable Bundle savedInstanceState) {
    AndroidInjection.inject(this);
    super.onCreate(savedInstanceState);
  }

  @Override
  protected void onResume() {
    super.onResume();
    mDisposables.resume();
  }

  @Override
  protected void onPause() {
    mDisposables.pause();
    super.onPause();
  }

  @Override
  protected void onDestroy() {
    mDisposables.dispose();
    super.onDestroy();
  }

  @Override
  public AndroidInjector<Fragment> fragmentInjector() {
    return mFragmentInjector;
  }

  protected void registerDisposable(Disposable disposable) {
    mDisposables.addDisposable(disposable);
  }

  protected void registerPausable(Pausable pausable) {
    mDisposables.addPausable(pausable);
  }
}
