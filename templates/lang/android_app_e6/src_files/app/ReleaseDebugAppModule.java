package {{ .Group }}.app;

import com.episode6.hackit.chop.Chop;
import com.episode6.hackit.chop.android.AndroidDebugTree;

import javax.inject.Singleton;

import dagger.Module;
import dagger.Provides;

/**
 * Release version of DebugAppModule
 */
@Module
public class DebugAppModule {

  @Provides
  @Singleton
  static Chop.Tree provideChopTree() {
    return new AndroidDebugTree("{{ .CamelNameWithoutApp }}App") {
      @Override
      public boolean supportsLevel(Chop.Level level) {
        return level == Chop.Level.E;
      }
    };
  }
}
