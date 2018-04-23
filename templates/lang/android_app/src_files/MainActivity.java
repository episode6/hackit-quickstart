package {{ .Group }};

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;

public class MainActivity extends AppCompatActivity {

  @Override
  protected void onCreate(Bundle savedInstanceState) {
    super.onCreate(savedInstanceState);
    setContentView(R.layout.activity_main);

    if (getFragmentManager().findFragmentById(R.id.main_content) == null) {
      getFragmentManager().beginTransaction()
          .add(R.id.main_content, new MainFragment())
          .commit();
    }
  }
}
