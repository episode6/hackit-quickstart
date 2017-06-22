package {{ .Group }}.base;

import com.episode6.hackit.android.inject.context.qualifier.ForFragment;
import com.episode6.hackit.android.inject.context.scope.FragmentScope;
import com.episode6.hackit.android.inject.thread.ForUiThread;
import com.episode6.hackit.disposable.DisposableManager;
import com.episode6.hackit.disposable.Disposables;
import com.episode6.hackit.pausable.PausableDisposableManager;
import com.episode6.hackit.pausable.PausableExecutor;
import com.episode6.hackit.pausable.PausableManager;
import com.episode6.hackit.pausable.Pausables;

import java.util.concurrent.Executor;

import javax.inject.Named;

import dagger.Binds;
import dagger.Module;
import dagger.Provides;

@Module
public abstract class BaseFragmentModule {

  @Provides
  static @Named("forFragmentUi") DisposableManager provideUiDisposableManager() {
    return Disposables.newManager();
  }

  @Provides
  @FragmentScope
  static @ForFragment PausableDisposableManager providePausableDisposableManager() {
    return Pausables.newDisposableManager();
  }

  @Binds
  @FragmentScope
  abstract @ForFragment DisposableManager bindScopedDisposableManager(
      @ForFragment PausableDisposableManager mgr);

  @Binds
  @FragmentScope
  abstract @ForFragment PausableManager bindScopedPausableManager(
      @ForFragment PausableDisposableManager mgr);

  @Provides
  @FragmentScope
  static @ForUiThread PausableExecutor bindPausableExecutor(
      @ForUiThread Executor uiExecutor,
      @ForFragment PausableManager pausableManager) {
    PausableExecutor pausableExecutor = Pausables.queuingExecutor(uiExecutor);
    pausableManager.addPausable(pausableExecutor);
    return pausableExecutor;
  }
}
