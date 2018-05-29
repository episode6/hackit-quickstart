package {{ .Group }}.main;

import android.os.Bundle;

import {{ .Group }}.R;
import {{ .Group }}.base.BaseAppCompatActivity;

public class MainActivity extends BaseAppCompatActivity {

  @Override
  protected void onCreate(Bundle savedInstanceState) {
    super.onCreate(savedInstanceState);
    setContentView(R.layout.activity_main);

    if (getSupportFragmentManager().findFragmentById(R.id.main_content) == null) {
      getSupportFragmentManager().beginTransaction()
          .add(R.id.main_content, new MainFragment())
          .commit();
    }
  }
}
