package {{ .Group }}.main;

import android.app.Activity;

import {{ .Group }}.base.BaseActivityModule;
import {{ .Group }}.base.BaseFragmentModule;

import com.episode6.hackit.android.inject.context.scope.ActivityScope;
import com.episode6.hackit.android.inject.context.scope.FragmentScope;

import dagger.Binds;
import dagger.Module;
import dagger.android.ContributesAndroidInjector;

@Module(includes = {BaseActivityModule.class})
public abstract class MainActivityModule {

  @Binds
  @ActivityScope
  abstract Activity bindMainActivity(MainActivity activity);

  @FragmentScope
  @ContributesAndroidInjector(modules = {BaseFragmentModule.class})
  abstract MainFragment mainFragment();
}
