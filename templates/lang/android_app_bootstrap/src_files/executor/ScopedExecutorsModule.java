package {{ .Group }}.executor;

import android.content.Context;
import android.os.Handler;
import android.os.HandlerThread;
import android.os.Process;
import android.support.annotation.NonNull;

import com.episode6.hackit.android.inject.context.qualifier.ForContextScope;
import com.episode6.hackit.android.inject.context.scope.ContextScope;
import com.episode6.hackit.android.inject.thread.ForWorkerThread;
import com.episode6.hackit.disposable.DisposableManager;
import com.episode6.hackit.disposable.android.AndroidDisposables;

import java.util.concurrent.Executor;

import dagger.Module;
import dagger.Provides;

@Module
public class ScopedExecutorsModule {

  @Provides
  @ContextScope
  static @ForWorkerThread HandlerThread provideBackgroundHandlerThread(
      Context context,
      @ForContextScope DisposableManager disposableManager) {
    HandlerThread workerThread = new HandlerThread(
        "WorkerFor:" + context.toString(),
        Process.THREAD_PRIORITY_BACKGROUND);
    workerThread.start();
    disposableManager.addDisposable(AndroidDisposables.forHandlerThread(workerThread));
    return workerThread;
  }

  @Provides
  @ContextScope
  static @ForWorkerThread Handler provideBackgroundHandler(@ForWorkerThread HandlerThread thread) {
    return new Handler(thread.getLooper());
  }

  @Provides
  @ContextScope
  static @ForWorkerThread Executor provideBackgroundExecutor(final @ForWorkerThread Handler handler) {
    return new Executor() {
      @Override
      public void execute(@NonNull Runnable command) {
        handler.post(command);
      }
    };
  }
}
