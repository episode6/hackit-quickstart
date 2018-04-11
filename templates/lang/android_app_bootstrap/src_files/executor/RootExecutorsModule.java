package {{ .Group }}.executor;

import android.os.Handler;
import android.os.Looper;
import android.support.annotation.NonNull;

import com.episode6.hackit.android.inject.thread.ForUiThread;

import java.util.concurrent.Executor;

import javax.inject.Singleton;

import dagger.Module;
import dagger.Provides;

@Module
public class RootExecutorsModule {

  @Provides
  @Singleton
  static @ForUiThread Handler provideUiHandler() {
    return new Handler(Looper.getMainLooper());
  }

  @Provides
  @Singleton
  static @ForUiThread Executor provideUiExecutor(final @ForUiThread Handler handler) {
    return new Executor() {
      @Override
      public void execute(@NonNull Runnable command) {
        if (Looper.myLooper() == Looper.getMainLooper()) {
          command.run();
        } else {
          handler.post(command);
        }
      }
    };
  }
}
