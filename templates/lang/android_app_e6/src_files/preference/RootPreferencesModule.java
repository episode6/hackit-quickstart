package {{ .Group }}.preference;

import android.content.Context;
import android.content.SharedPreferences;
import android.preference.PreferenceManager;

import com.episode6.hackit.android.inject.context.qualifier.ForApplication;
import com.episode6.hackit.typed.preferences.TypedPrefs;

import javax.inject.Singleton;

import dagger.Module;
import dagger.Provides;

@Module
public class RootPreferencesModule {

  @Provides
  @Singleton
  static SharedPreferences provideDefaultSharedPrefs(@ForApplication Context applicationContext) {
    return PreferenceManager.getDefaultSharedPreferences(applicationContext);
  }

  @Provides
  @Singleton
  static TypedPrefs provideDefaultTypedPrefs(SharedPreferences sharedPreferences) {
    return TypedPrefs.Wrap.sharedPrefs(sharedPreferences);
  }
}
