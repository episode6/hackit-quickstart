package {{ .Group }}.app;

import com.episode6.hackit.chop.Chop;
import com.episode6.hackit.chop.android.AndroidDebugTree;

import javax.inject.Singleton;

import dagger.Module;
import dagger.Provides;

/**
 * Debug version of DebugAppModule
 */
@Module
public class DebugAppModule {

  @Provides
  @Singleton
  static Chop.Tree provideChopTree() {
    return new AndroidDebugTree("{{ .CamelNameWithoutApp }}App");
  }
}
