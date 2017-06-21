package {{ .Group }}.base;

import android.app.Activity;
import android.content.Context;

import com.episode6.hackit.android.inject.context.module.ScopedContextModule;
import com.episode6.hackit.android.inject.context.qualifier.ForContextScope;
import com.episode6.hackit.android.inject.context.scope.ContextScope;
import com.episode6.hackit.disposable.DisposableManager;
import com.episode6.hackit.pausable.PausableDisposableManager;
import com.episode6.hackit.pausable.PausableManager;
import com.episode6.hackit.pausable.Pausables;

import dagger.Binds;
import dagger.Module;
import dagger.Provides;

@Module(includes = {ScopedContextModule.class})
public abstract class BaseActivityModule {

  @Binds
  @ContextScope
  abstract Context bindContext(Activity activity);

  @Provides
  @ContextScope
  static @ForContextScope PausableDisposableManager providePausableDisposableManager() {
    return Pausables.newDisposableManager();
  }

  @Binds
  @ContextScope
  abstract @ForContextScope DisposableManager bindScopedDisposableManager(
      @ForContextScope PausableDisposableManager mgr);

  @Binds
  @ContextScope
  abstract @ForContextScope PausableManager bindScopedPausableManager(
      @ForContextScope PausableDisposableManager mgr);
}
